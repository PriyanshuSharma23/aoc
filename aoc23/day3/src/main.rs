use std::{fs, io::Read};

fn main() {
    let mut file = fs::File::open("input.txt").unwrap();
    let mut data = String::new();
    let _ = file.read_to_string(&mut data);

    let output = solve(&data);
    println!("{output}");
}

fn solve(data: &str) -> i32 {
    let grid: Vec<Vec<char>> = data.lines().map(|s| s.chars().collect()).collect();
    let h = grid.len();
    let w = grid[0].len();

    let mut builder = 0;
    let mut char_found = false;
    let mut total = 0;
    for i in 0..h {
        for j in 0..w {
            if grid[i][j].is_digit(10) {
                let ch = grid[i][j] as i32 - '0' as i32;
                builder = builder * 10 + ch;

                if i > 0 {
                    if j > 0 {
                        let tl = grid[i - 1][j - 1];
                        if check_symbol(tl) {
                            char_found = true;
                        }
                        if check_symbol(grid[i][j - 1]) {
                            char_found = true;
                        }
                    }

                    if j + 1 < w {
                        let tr = grid[i - 1][j + 1];
                        if check_symbol(tr) {
                            char_found = true;
                        }
                        if check_symbol(grid[i][j + 1]) {
                            char_found = true;
                        }
                    }
                }

                if i + 1 < h {
                    if j > 0 {
                        let bl = grid[i + 1][j - 1];
                        if check_symbol(bl) {
                            char_found = true;
                        }
                        if check_symbol(grid[i][j - 1]) {
                            char_found = true;
                        }
                    }

                    if j + 1 < w {
                        let br = grid[i + 1][j + 1];
                        if check_symbol(br) {
                            char_found = true;
                        }
                        if check_symbol(grid[i][j + 1]) {
                            char_found = true;
                        }
                    }
                }
            } else {
                if char_found {
                    total += builder;
                }
                if builder > 0 && !char_found {
                    println!("builder: {builder}");
                    if i > 0 {
                        println!("{}", grid[i - 1].iter().collect::<String>());
                    }
                    println!("{}", grid[i].iter().collect::<String>());
                    if i < h - 1 {
                        println!("{}", grid[i + 1].iter().collect::<String>());
                    }

                    println!("\n\n");
                }
                char_found = false;
                builder = 0;
            }
        }

        if char_found {
            total += builder;
            builder = 0;
            char_found = false;
        }
    }

    return total;
}

fn check_symbol(c: char) -> bool {
    // return c == '#' || c == '@' || c == '$' || c == '%' || c == '&' || c == '*';
    return !c.is_digit(10) && c != '.';
}
