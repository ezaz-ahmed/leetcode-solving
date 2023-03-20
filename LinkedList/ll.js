class Node {
  constructor(value) {
    this.value = value;
    this.next = null;
  }
}

class LinkedList {
  constructor(value) {
    const node = new Node(value);
    this.head = node;
    this.tail = this.head;
    this.length = 1;
  }

  printLinkedList() {
    let temp = this.head;
    while (temp !== null) {
      console.log(this.head.value);
      temp = temp.next;
    }
  }
}

const test = () => {
  const ll = new LinkedList(1);
  ll.printLinkedList();
};

test();
