import React from 'react'

class AddUserBox extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      backgroundImg: "",
    };
  }

  addUser() {
    console.log("Adding new user.");
    this.props.onUserAdded({
      userId: 'blabla',
      name: this.state.name,
      profileImg: this.state.backgroundImg,
      pos: {
        x: 1,
        y: 1,
      },
    })
  }

  handleNameChange = (e) => {
    this.setState({name: e.target.value});
  }

  handleProfileImgSelected= (e) => {
    console.log(e.target.getAttribute("img"));
    this.setState({backgroundImg: e.target.getAttribute("img")});
  }

  profileImgStyle(imgUrl) {
    return {
      backgroundImage: 'url("' + imgUrl + '")',
      height: '70px',
      width: '70px',
      backgroundSize: 'auto 100%',
      display: 'inline-block',
    }
  }

  renderProfileImg(imgUrl) {
    return (
          <div style={this.profileImgStyle(imgUrl)} img={imgUrl} onClick={this.handleProfileImgSelected}>
          </div>
        );
  }

  render() {
    const style = {
      width: '100%',
      height: '100%',
      backgroundColor: '#8e87879c',
      display: 'inline-block',
    };

    return (
      <div style={style}>
        <div>
          {this.renderProfileImg('/avatar-1.png')}
          {this.renderProfileImg('/avatar-2.png')}
          {this.renderProfileImg('/avatar-3.png')}
        </div>
        <input type="text" value={this.state.value} onChange={this.handleNameChange}/> 
        <button onClick={this.addUser.bind(this)}>Add</button>
        ---

        {this.state.name}, {this.state.backgroundImg}
      </div>
    )
  }
}

export default AddUserBox;
