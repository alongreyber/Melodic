export class Message {
    text: String;
    color: String;
    constructor(text: String, color: string) {
	this.text = text;
	this.color = color; 
    }
}

class Store {
  readonly messages: Array<Message> = [];

  addMessage (m: Message) {
      this.messages.push(m)
  };

  removeMessageByIndex(index: number) {
    this.messages.splice(index, 1);
  };

}

export var store = new Store();
