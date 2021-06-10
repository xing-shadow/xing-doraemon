import React, { Component } from 'react';
import {BrowserRouter,Switch,Redirect,Route} from 'react-router-dom';
import Index from './view/layout/Index';
import Login from './view/login/Index';
class App extends Component {
  constructor(props) {
    super(props);
    this.state = {  }
  }
  render() { 
    return (
        <BrowserRouter>
          <Switch>
            <Route component={Index} path="/antd/dist/"></Route>
            <Route exact component={Login} path="/antd/login/"></Route>
            <Redirect exact from="/" to="/antd/dist/" />
          </Switch>
        </BrowserRouter>
    );
  }
}
 
export default App;
