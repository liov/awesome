use pyo3::prelude::*;
use pyo3::types::PyModule;
use std::io::Read;
use std::path::Path;
use tendril::stream::TendrilSink;

/// Formats the sum of two numbers as string.
#[pyfunction]
fn sum_as_string(a: usize, b: usize) -> PyResult<String> {
    Ok((a + b).to_string())
}

/// A parsed html document
#[pyclass(unsendable)]
struct Document {
    node: kuchiki::NodeRef,
}

#[pymethods]
impl Document {
    /// Returns the selected elements as strings
    fn select(&self, selector: &str) -> Vec<String> {
        self.node
            .select(selector)
            .unwrap()
            .map(|css_match| css_match.text_contents())
            .collect()
    }
}

impl Document {
    fn from_reader(reader: &mut impl Read) -> PyResult<Document> {
        let node = kuchiki::parse_html().from_utf8().read_from(reader)?;
        Ok(Document { node })
    }

    fn from_file(path: &Path) -> PyResult<Document> {
        let node = kuchiki::parse_html().from_utf8().from_file(path)?;
        Ok(Document { node })
    }
}

/// Parses the File from the specified Path into a document
#[pyfunction]
fn parse_file(path: &str) -> PyResult<Document> {
    let document = Document::from_file(Path::new(path))?;
    Ok(document)
}

/// Parses the given html test into a document
#[pyfunction]
fn parse_text(text: &str) -> PyResult<Document> {
    let document = Document::from_reader(&mut text.as_bytes())?;
    Ok(document)
}

#[pymodule]
fn _lib(m: &Bound<'_, PyModule>) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_as_string, m)?)?;
    m.add_class::<Document>()?;
    m.add_function(wrap_pyfunction!(parse_file, m)?)?;
    m.add_function(wrap_pyfunction!(parse_text, m)?)?;
    Ok(())
}
