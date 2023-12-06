use std::{fs, io::Read};

fn main() {
    let mut file = fs::File::open("input.txt").unwrap();
    let mut data = String::new();
    let _ = file.read_to_string(&mut data);

    let output = solve(&data);
    println!("{output}");
}

fn solve(input: &str) -> u32 {
    let mut words: Vec<&str> = vec![
        "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let mut nums: Vec<&str> = vec!["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"];

    return input
        .lines()
        .map(|line| {
            let mut number: u32 = 0;
            let mut least_front = line.len();
            let mut max_last = -1;

            let mut closure =
                |comp: &mut Vec<&str>, least_front: &mut usize, max_last: &mut i32| {
                    for (idx, word) in comp.iter().enumerate() {
                        if let Some(i) = line.find(word) {
                            match (*least_front).cmp(&i) {
                                std::cmp::Ordering::Greater => {
                                    *least_front = i;
                                    number = (idx as u32) * 10 + (number % 10);
                                }
                                _ => (),
                            }
                        }

                        if let Some(i) = line.rfind(word) {
                            match (*max_last).cmp(&(i as i32)) {
                                std::cmp::Ordering::Less => {
                                    *max_last = i as i32;
                                    number = (number / 10) * 10 + idx as u32;
                                }
                                _ => (),
                            }
                        }
                    }
                };

            closure(&mut words, &mut least_front, &mut max_last);
            closure(&mut nums, &mut least_front, &mut max_last);

            println!("{number}");

            return number;
        })
        .sum();
}
