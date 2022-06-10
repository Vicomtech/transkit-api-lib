import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="Transkit",
    packages = ['Transkit'],
    version="1.0.0",
    author="Vicomtech",
    author_email="posdig-dev@vicomtech.org",
    description="Transkit API library",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/Vicomtech/transkit-api-lib/",
    # packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: GNU General Public License v3 (GPLv3)",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
)
