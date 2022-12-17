use std::collections::HashMap;

pub struct Parser {
    chs: Vec<char>,
    pos: usize,
}

impl Parser {
    pub fn new(code: &str) -> Self {
        Parser {
            chs: code.chars().collect::<Vec<char>>(),
            pos: 0,
        }
    }

    pub fn parse<'a>(&mut self) -> Result<HashMap<(String, char), (String, char, char)>, &'a str> {
        let mut rules = HashMap::new();
        while self.pos < self.chs.len() {
            if self.chs[self.pos].is_whitespace() {
                self.pos += 1;
                continue;
            }

            self.consume('(')?;

            let cur_state = self.read_str();

            self.consume(',')?;

            let input = self.read_char()?;

            self.consume(')')?;
            self.consume(':')?;
            self.consume('(')?;

            let next_state = self.read_str();

            self.consume(',')?;

            let output = self.read_char()?;

            self.consume(',')?;

            let dir = self.read_char()?;

            rules.insert(
                (cur_state, input),
                (next_state, output, dir)
            );

            self.pos += 1;
        }
        // println!("{:?}", rules);
        return Ok(rules);
    }

    fn read_char<'a>(&mut self) -> Result<char, &'a str> {
        if self.pos < self.chs.len() {
            let ch = self.chs[self.pos];
            self.pos += 1;
            return Ok(ch);
        } else {
            return Err("Syntax Error");
        }
    }

    fn read_str(&mut self) -> String {
        let mut s = String::from("");
        while self.pos < self.chs.len() {
            if self.chs[self.pos] == ',' || self.chs[self.pos] == '(' || self.chs[self.pos] == ')' || self.chs[self.pos] == '\n' {
                return s;
            }
            s.push(self.chs[self.pos]);
            self.pos += 1;
        }
        return s;
    }

    fn consume<'a>(&mut self, ch: char) -> Result<(), &'a str> {
        if self.chs[self.pos] == ch {
            self.pos += 1;
            return Ok(());
        } else {
            return Err("Syntax Error");
        }
    }
}

pub struct Turing {
    pos: usize,
    state: String,
    rules: HashMap<(String, char), (String, char, char)>,
    tape: Vec<char>
}

impl Turing {
    pub fn new<'a>(code: &str) -> Result<Self, &'a str> {
        match Parser::new(code).parse() {
            Ok(rules) => Ok(Turing {
                pos: 0,
                state: String::from("S"),
                rules: rules,
                tape: Vec::<char>::new(),
            }),
            Err(msg) => Err(msg),
        }
    }

    pub fn set_tape(&mut self, tape: &str) {
        self.tape = tape.chars().collect::<Vec<char>>();
    }

    pub fn run(&mut self) -> Vec<char> {
        while self.state != "_" {
            self.step(); // println!("{} {:?}", self.state, self.tape);
        }
        return self.tape.clone();
    }
    
    pub fn step(&mut self) -> Vec<char> {
        let input = self.tape[self.pos];
        if let Some(t) = self.rules.get(&(self.state.clone(), input)).clone() {
            let (next_state, output, dir) = (*t).clone();
            self.state = next_state;
            self.tape[self.pos] = output;
            if dir == 'L' {
                self.pos -= 1;
            } else if dir == 'R' {
                self.pos += 1;
            }
        } else {
            self.state = String::from("_");
        }

        return self.tape.clone();
    }
}