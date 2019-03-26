from setuptools import setup

setup(
  name="mydsl",
  version="0.0.3",
  install_requires=["aiohttp", "pymongo", "pyyaml", "jinja2"],
  extras_require={
  },
  entry_points={
      "console_scripts": [
      ],
      "gui_scripts": [
      ]
  }
)
