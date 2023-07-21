from setuptools import setup, find_packages

setup(
  name='am4',
  packages=['am4'],
  version='1.0',
  license='MIT',
  description='Python bot for Airline Manager 4 game',
  author='Andrei Shokhin',
  author_email='shokhin.andrey@gmail.com',
  url='https://github.com/ashokhin/am4b',
  keywords=['airlinemanager', 'bot', 'game', 'python'],
  packages=find_packages(),
  install_requires=['selenium'],

  classifiers=[
    'Development Status :: 3 - Alpha',
    'Intended Audience :: Developers',
    'Topic :: Software Development :: Build Tools',
    'License :: OSI Approved :: MIT License',
    'Programming Language :: Python :: 3',
  ],
)