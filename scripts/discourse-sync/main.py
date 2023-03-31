import os
import sys
import yaml
from pydiscourse import DiscourseClient


def main():
    # Get configuration from environment variables
    DISCOURSE_HOST = os.environ.get('DISCOURSE_HOST', 'https://discourse.charmhub.io/')
    DISCOURSE_API_USERNAME = os.environ.get('DISCOURSE_API_USERNAME')
    DISCOURSE_API_KEY = os.environ.get('DISCOURSE_API_KEY')
    DOCS_DIR = os.environ.get('DOCS_DIR')
    TOPIC_IDS = os.environ.get('TOPIC_IDS')
    CI = os.environ.get('CI')

    client = DiscourseClient(
        host=DISCOURSE_HOST,
        api_username=DISCOURSE_API_USERNAME,
        api_key=DISCOURSE_API_KEY,
    )

    with open(TOPIC_IDS, 'r') as file:
        topic_ids = yaml.safe_load(file)
        if topic_ids is None:
            topic_ids = {}

    # Keep track of docs which don't have a matching topic ID.
    # If any are found while running in CI, we need to fail and ask the
    # developer to manually intervene.
    not_found_docs = []

    for entry in os.scandir(DOCS_DIR):
        if not is_markdown_file(entry):
            print(f'file {entry.name}: not a Markdown file: skipping')
            continue

        doc_name = removesuffix(entry.name, ".md")
        content = open(entry.path, 'r').read()

        if topic_ids and doc_name in topic_ids:
            print(f'doc {doc_name} (topic #{topic_ids[doc_name]}): checking for changes')
            # API call to get the post ID from the topic ID
            # TODO: we could save the post IDs in a separate yaml file and
            #   avoid this extra API call
            topic = client.topic(
                slug='',
                topic_id=topic_ids[doc_name],
            )
            post_id = topic['post_stream']['posts'][0]['id']

            # Get current contents of post
            post2 = client.post_by_id(
                post_id=post_id
            )
            current_contents = post2['raw']
            if current_contents == content.rstrip('\n'):
                print(f'doc {doc_name} (topic #{topic_ids[doc_name]}): already up-to-date: skipping')
                continue

            # Update Discourse post
            print(f'doc {doc_name} (topic #{topic_ids[doc_name]}): updating')
            client.update_post(
                post_id=post_id,
                content=content,
            )

        elif CI == 'true':
            # Running in CI - we can't edit the TOPIC_IDS yaml file.
            # Log a warning and return a non-zero exit code later
            print(f'WARNING: no topic ID found for doc {doc_name}')
            not_found_docs.append(doc_name)

        else:
            # Create new Discourse post
            print(f'doc {doc_name}: no topic ID found: creating new post')
            post = client.create_post(
                title=post_title(doc_name),
                category_id=22,
                content=content,
                tags=['olm', 'autogenerated'],
            )
            new_topic_id = post['topic_id']
            print(f'doc {doc_name}: created new topic #{new_topic_id}')

            # Save topic ID in yaml map for later
            topic_ids[doc_name] = new_topic_id
            with open(TOPIC_IDS, 'w') as file:
                yaml.safe_dump(topic_ids, file)

    if len(not_found_docs) > 0:
        print("The following docs don't have corresponding entries in the TOPIC_IDS yaml file. Please run "
              "scripts/discourse-sync locally to create new topics and update the topic IDs.")
        for doc_name in not_found_docs:
            print(f' - {doc_name}')
        sys.exit(1)


def is_markdown_file(entry: os.DirEntry) -> bool:
    return entry.is_file() and entry.name.endswith(".md")


def removesuffix(text, suffix):
    if suffix and text.endswith(suffix):
        return text[:-len(suffix)]
    return text


def post_title(doc_name: str) -> str:
    if doc_name == 'index':
        return 'Juju CLI commands'
    return f"Command '{doc_name}'"


if __name__ == "__main__":
    main()
