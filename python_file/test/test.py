import addressbook_pb2


def create_address_book():
    address_book = addressbook_pb2.AddressBook()

    person = address_book.people.add()
    person.name = "John Doe"
    person.id = 123
    person.email = "johndoe@example.com"

    phone = person.phones.add()
    phone.number = "555-1234"
    phone.type = addressbook_pb2.Person.HOME

    return address_book


def write_address_book(address_book, filename):
    with open(filename, "wb") as f:
        f.write(address_book.SerializeToString())


def read_address_book(filename):
    address_book = addressbook_pb2.AddressBook()

    with open(filename, "rb") as f:
        address_book.ParseFromString(f.read())

    return address_book


# 创建地址簿
address_book = create_address_book()

# 将地址簿写入二进制文件
write_address_book(address_book, "address_book.bin")

# 从二进制文件中读取地址簿
read_address_book = read_address_book("address_book.bin")

# 打印读取的地址簿数据
print(read_address_book)
