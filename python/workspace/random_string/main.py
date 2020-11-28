import random
import string


def random_string(count: int) -> str:
    return "".join([random.choice(string.ascii_letters) for _ in range(count)])


if __name__ == "__main__":
    print(random_string(6))
    print(random_string(10))
    print(random_string(20))
