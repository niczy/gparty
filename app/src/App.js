import React from 'react';
import logo from './logo.svg';
import './App.css';

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
    this.userState = props.userState;
  }

  getGridStyle(userState) {
    if (userState) {
      return {
        backgroundImage: 'url("' + userState.backgroundImg + '")',
        backgroundSize: 'auto 100%'
      };
    } else {
      return {}
    }
  }

  render() {
    return (
      <div className="grid" style={this.getGridStyle(this.userState)}></div>
    )
  }
}

class Board extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      desiredGrids: [],
      currentGrids: [
        [{backgroundImg: "/avatar-1.png"}, {}, {}],
        [{}, {}, {backgroundImg: "/avatar-2.png"}],
        [{}, {backgroundImg: "/avatar-3.png"}, {}],
      ]
    }
  }

  renderRow(row, idx) {
    return (
        <div key={idx}>
          { row.map((grid, gIdx) => <Grid  key={idx + '-' + gIdx} userState={grid} />)}
        </div>
    )
  }

  render() {
    return (
        <div>
          {
            this.state.currentGrids.map((row, idx) => this.renderRow(row, idx))
          }
        </div>
        );
  }
}

export default App;
