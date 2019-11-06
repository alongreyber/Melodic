import Message from "./message.ts";

class Store {
  public readonly messages: Message[] = [];

  public addMessage(m: Message) {
      this.messages.push(m)
  }

  public removeMessageByIndex(index: number) {
    this.messages.splice(index, 1);
  }

}

export var store = new Store();
