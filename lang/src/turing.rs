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

    pub fn parse(&mut self) -> Result<HashMap<(&str, char), (&str, char, char)>, &str> {
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
        return Ok(rules);
    }

    fn read_char(&mut self) -> Result<char, &str> {
        if self.pos < self.chs.len() {
            let ch = self.chs[self.pos];
            self.pos += 1;
            return Ok(ch);
        } else {
            return Err("Syntax Error");
        }
    }

    fn read_str(&mut self) -> &str {
        let mut s = String::from("");
        while self.pos < self.chs.len() {
            if self.chs[self.pos] == ',' || self.chs[self.pos] == '(' || self.chs[self.pos] == ')' || self.chs[self.pos] == '\n' {
                self.pos += 1;
                return &*s;
            }
            s.push(self.chs[self.pos]);
            self.pos += 1;
        }
        return &*s;
    }

    fn consume(&mut self, ch: char) -> Result<(), &str> {
        if self.chs[self.pos] == ch {
            self.pos += 1;
            return Ok(());
        } else {
            return Err("Syntax Error");
        }
    }
}

pub struct Turing<'a> {
    pos: usize,
    state: &'a str,
    rules: HashMap<(&'a str, char), (&'a str, char, char)>,
}

impl Turing<'_> {
    pub fn new(code: &str) -> Result<Self, &str> {
        match Parser::new(code).parse() {
            Ok(rules) => Ok(Turing {
                pos: 0,
                state: "S",
                rules: rules
            }),
            Err(msg) => Err(msg),
        }
    }

    pub fn compute(&self, tape: &str) -> Vec<char> {
        let tape_chars: Vec<char> = tape.chars().collect();

        while self.state != "_" {
            let input = tape_chars[self.pos];
            
            if let Some(&(next_state, output, dir)) = self.rules.get(&(self.state, input)) {
                self.state = next_state;
                tape_chars[self.pos] = output;
                if dir == 'L' {
                    self.pos -= 1;
                } else if dir == 'R' {
                    self.pos += 1;
                }
            } else {
                return tape_chars;
            }
        }
        return tape_chars;
    }
}