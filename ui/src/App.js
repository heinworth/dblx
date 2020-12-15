import './App.css';
import React, { Component } from 'react'

class App extends Component {

  url = 'http://localhost:8080/'
  state = {
    users: [],
    selectedId: 0,
    selectedIndex: -1,
    nameInput: ''
  }

  componentDidMount() {
    fetch(this.url + `user`)
    .then(res => res.json())
    .then((data) => {
      this.setState({users: [...data]})
      this.render()
    })
    .catch(console.log)
  }


  addUser() {
    fetch(this.url + `user/${this.state.nameInput}`, {method: 'POST'})
    .then(res => res.json())
    .then((data) => {
      this.setState({users: [...this.state.users, data]})
      this.render()
    })
    .catch(console.log)
  }

  userClicked(i) {
    this.setState({selectedIndex: i})
    this.getUser(i)
  }

  getUser(i) {
    let id = this.state.users[i].id
    fetch(this.url + `user/${id}`)
    .then(res => res.json())
    .then((data) => {
      this.setState({selectedId: data})
      this.render()
    })
    .catch(console.log)
  }

  nameChanged(event) {
    const name = event.target.value
    this.setState({nameInput: name})
  }

  render() {
    return (
      <div class="app columns">
        <div class="col">
          <form>
            <p class="form-label">Enter user name:</p>
            <input
              type="text"
              class="input"
              onChange={(event) => this.nameChanged(event)}
            />
          </form>
          <button class="add-button" onClick={() => this.addUser()}>
            Add
          </button>
        </div>
        <div class="col">
          <u> Users:</u> 
          <div class="small"> 
            (click to show ID) 
          </div>
          <br/>
            {this.state?.users.map((user, i) => {
              return <div 
                style={{color: this.state.selectedIndex === i ? "yellow" : "white"}}
                onClick={() => this.userClicked(i)}
                key={i}> 
                { this.state.selectedIndex === i ? user.id : user.name} 
              </div>
            })} 
          </div>
      </div>
    )
  }
}

export default App;
