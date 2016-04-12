#ifndef __LINKED_LIST_H__
#define __LINKED_LIST_H__

#include <atomic>

using namespace std;

class Node {
  friend class LinkedList;
private:
  atomic<int> item;
  atomic<int> key;
  atomic<Node> *next;
public:
  Node(int val) {
    item.Store(val);
    key = val.hashCode();
    next.Set(NULL);
  }

  int Key(void) {
    return key.Load();
  }
};

class LinkedList {
private:
  atomic<Node> *head;
  atomic<Node> *tail;
public:
  bool add(int item) {

  }

  bool remove(int item) {
    int key = item.hashCode();
    bool snip;

    while(true) {

    }
  }

  bool contains(int item) {

  }

};

#endif
