pub use crate::vscode::open;

pub fn read(path: &str) {
    println!("This is vscode open command.");
    open::get_yml(path);
}
