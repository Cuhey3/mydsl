from setuptools import setup

setup(
  name="mydsl",
  version="0.0.1",
  install_requires=["aiohttp", "pymongo"],
  extras_require={
  },
  entry_points={
      "console_scripts": [
      ],
      "gui_scripts": [
      ]
  }
)
