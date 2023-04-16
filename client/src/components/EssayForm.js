import * as React from 'react';
import axios from 'axios';

export default class EssayForm extends React.Component {
    constructor(props) {
      super(props);
      this.state = {
        value: 'めもだよっ'
      };
  
      this.handleChange = this.handleChange.bind(this);
      this.handleSubmit = this.handleSubmit.bind(this);
    }
  
    handleChange(event) {
      this.setState({value: event.target.value});
    }
  
    handleSubmit(event){
      
      let params = new URLSearchParams();
      let text  = this.state.value;
      params.append('apikey', text);
      axios.post("http://127.0.0.1:8080/post", params)
        .then(function (response) {
          // 送信成功時の処理
          console.log(response);
          alert('メモが保存されました: ' + text);
        })
        .catch(function (error) {
          // 送信失敗時の処理
          console.log(error);
          alert(error);
        });
        
      event.preventDefault();
     
    }
  
    render() {
      return (
        <form onSubmit={this.handleSubmit}>
          <label>
            <textarea value={this.state.value} onChange={this.handleChange} />
          </label>
          <input type="submit" value="Submit" />
        </form>
      );
    }
  }
  