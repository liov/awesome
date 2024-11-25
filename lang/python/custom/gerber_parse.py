
from pygerber.gerberx3.parser2.parser2 import Parser2

from pygerber.gerberx3.tokenizer.tokenizer import Tokenizer
with open(r"xxx", 'r', encoding='utf-8') as file:
    # 读取文件内容
    content = file.read()
    stack = Tokenizer().tokenize(content)
    cmd_buf = Parser2().parse(stack)
    print(cmd_buf)
