export const store = {
  state: {
    messages: []
  },
  addMessage (newMessage) {
    this.state.messages.push(newMessage)
  },
  removeMessageByIndex(index) {
    this.state.messages.splice(index, 1);
  }
}
