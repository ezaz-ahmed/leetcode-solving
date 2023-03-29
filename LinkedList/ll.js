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

  getHead() {
    if (this.head === null) {
      console.log('Head: null');
    } else {
      console.log(`Head: ${this.head.value}`);
    }
  }

  getTail() {
    if (this.tail === null) {
      console.log('Tail: null');
    } else {
      console.log(`Tail: ${this.tail.value}`);
    }
  }

  getLength() {
    console.log('Length: ' + this.length);
  }

  makeEmpty() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }

  push(value) {
    const node = new Node(value);
    if (!this.head) {
      this.head = node;
      this.tail = node;
    }

    this.tail.next = node;
    this.tail = node;
    this.length++;

    return this;
  }
}

const test = () => {
  let ll = new LinkedList(1);
  ll.push(2);
  ll.push(3);
  ll.push(4);

  ll.printLinkedList();

  ll.getLength();
};

test();
