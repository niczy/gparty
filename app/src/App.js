import React from 'react';
import logo from './logo.svg';
import AddUserBox from './AddUserBox.js'
import './App.css';

const axios = require('axios');

function App() {
  return (
    <div className="App">
      <Board />
    </div>
  );
}

class Grid extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      userState: props.userState,
    };
  }

  getGridStyle(userState) {
    if (userState) {
      return {
        backgroundImage: 'url("' + userState.profileImg + '")',
        backgroundSize: 'auto 100%'
      };
    } else {
      return {}
    }
  }

  render() {
    return (
      <div className="grid" style={this.getGridStyle(this.state.userState)}></div>
    )
  }
}

const KEY_UP = 38;
const KEY_DOWN = 40;
const KEY_LEFT = 37;
const KEY_RIGHT = 39;
const MAX_X = 6;
const MAX_Y = 8;


class Board extends React.Component {
    constructor(props) {
    super(props);
    this.state = {
      currentUser: null,
      desiredGrids: [],
      currentGrids: [],
    }
    for (var i = 0; i < MAX_X; i++) {
      var row = [];
      for (var j = 0; j < MAX_Y; j++) {
        row.push({});
      }
      this.state.currentGrids.push(row);
    }

  }

  onUserAdded(user) {
    const x = user.pos.x;
    const y = user.pos.y;
    this.state.currentGrids[x][y] = user;
    this.setState({currentUser: user});
    // this.setState({currentGrids: this.state.currentGrids});
  }

  moveUser(delX, delY) {
    var x = this.state.currentUser.pos.x;
    var y = this.state.currentUser.pos.y;
    var newX = x + delX, newY = y + delY;
    if (!(newX >= 0 && newX < MAX_X && newY >= 0 && newY < MAX_Y)) {
      return;
    }
    this.state.currentGrids[x][y] = {};
    x += delX;
    y += delY;
    this.state.currentUser.pos.x = x;
    this.state.currentUser.pos.y = y;
    this.state.currentGrids[x][y] = this.state.currentUser;
    this.setState({currentGrids: this.state.currentGrids});
  }

  onKeyDown = (e) => {
    if (!this.state.currentUser) {
      return;
    }

    switch (e.keyCode) {
      case KEY_UP:
        e.preventDefault();
        this.moveUser(-1, 0);
        break;
      case KEY_DOWN:
        e.preventDefault();
        this.moveUser(1, 0);
        break;
      case KEY_LEFT:
        e.preventDefault();
        this.moveUser(0, -1);
        break;
      case KEY_RIGHT:
        e.preventDefault();
        this.moveUser(0, 1);
        break;
    }
  }

  componentDidMount() {
    axios.get('http://nicz.c.googlers.com:8060/_/getUserStates', {crossDomain: true})
        .then(res => {
          console.log(res);
        }).catch(error => {
          console.log(error);
        });
 
    document.addEventListener('keydown', this.onKeyDown)
  }

  componentWillUnmount() {
    document.removeEventListener('keydown', this.onKeyDown);
  }

  renderRow(row, idx) {
    return (
        <div key={idx} >
          { row.map((grid, gIdx) => <Grid  key={idx + '-' + gIdx + '-' + grid.userId} userState={grid} />)}
        </div>
    )
  }

  render() {
    if (!this.state.currentUser) {
      return (
          <AddUserBox onUserAdded={this.onUserAdded.bind(this)} />

      )
    } else {
      return (
          <div>
            {
              this.state.currentGrids.map((row, idx) => this.renderRow(row, idx))
            }
          </div>
      );
    }
  }
}

export default App;
