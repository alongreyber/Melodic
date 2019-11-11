import Message from "./message";

class Store {
  public readonly messages: Message[] = [];

  public addMessage(m: Message) {
      this.messages.push(m);
  }

  public removeMessageByIndex(index: number) {
    this.messages.splice(index, 1);
  }

}

const store = new Store();
export default store;
