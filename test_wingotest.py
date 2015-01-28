from mock import patch
import os
import shutil
import tarfile
from unittest import TestCase

from wingotest import (
    go_test_package,
    untar_gopath,
)
from utility import temp_dir


class WinGoTestTestCase(TestCase):

    def test_untar_gopath(self):
        with temp_dir() as base_dir:
            old_dir = os.path.join(base_dir, 'juju_1.2.3')
            for sub in ['bin', 'pkg', 'src']:
                os.makedirs(os.path.join(old_dir, sub))
            tarfile_path = os.path.join(base_dir, 'juju_1.2.3.tar.gz')
            with tarfile.open(name=tarfile_path, mode='w:gz') as tar:
                tar.add(old_dir, arcname='juju_1.2.3')
            shutil.rmtree(old_dir)
            gopath = os.path.join(base_dir, 'gogo')
            untar_gopath(tarfile_path, gopath, delete=True)
            self.assertTrue(os.path.isdir(gopath))
            self.assertEqual(['bin', 'src', 'pkg'], os.listdir(gopath))
            self.assertFalse(os.path.isfile(tarfile_path))

    def test_go_test_package(self):
        with temp_dir() as gopath:
            package_path = os.path.join(
                gopath, 'src', 'github', 'juju', 'juju')
            os.makedirs(package_path)
            with patch('wingotest.run', return_value=[0, 'success'],
                       autospec=True) as run_mock:
                devnull = open(os.devnull, 'w')
                with patch('sys.stdout', devnull):
                    return_code = go_test_package(
                        'github/juju/juju', 'go', gopath)
                    self.assertEqual(0, return_code)
                    args, kwargs = run_mock.call_args
                    self.assertEqual(('go', 'test', './...'), args)
                    self.assertEqual('amd64', kwargs['env'].get('GOARCH'))
                    self.assertEqual(gopath, kwargs['env'].get('GOPATH'))
