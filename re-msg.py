import re
import random
import string

msg_file = "./pkg/msg/msg.go"

with open(msg_file, 'r', encoding='utf-8') as f1, open(f"{msg_file}.bak", 'w+', encoding='utf-8') as f2:
    re1 = re.compile("json:\"(.*?)\"")
    for line in f1:
        ran_str = ''.join(random.sample(string.ascii_letters + string.digits, 8))
        f2.write(re.sub(re1, f"json:\"{ran_str}\"", line))
