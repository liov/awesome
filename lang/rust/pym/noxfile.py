import nox

@nox.session()
def test(session: nox.Session):
    session.install("setuptools", "setuptools-rust", "cffi")
    session.install("--no-build-isolation", ".")
    session.run("sum-cli", "1", "2")
    session.run("python", "-c", "import hello_world; print(hello_world)")