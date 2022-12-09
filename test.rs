

fn main() {
    let mut line = String::new();
    println!("Hello Dan! This is a test! If you see this message. We are good to go!");
    println!("Press enter to exit.");
    std::io::stdin().read_line(&mut line).unwrap();
    println!("Good by");
}