import React, { Component } from 'react';
import {Switch,Route} from 'react-router-dom';

const files = require.context("../../view/",true,/\.js$/)
const compoents = [];
files.keys().map(item => {
    if (item.includes("./layout") || item.includes("./login")) {
        return false;
    }
    let jsonObj = {};
    let splitFileName = item.split('.');
    let path = `/antd/dist${splitFileName[1].toLowerCase()}`
    let component = files(item).default;
    jsonObj.path = path;
    jsonObj.component = component;
    compoents.push(jsonObj);
    return false;
})

class ContaninerMain extends Component {
    constructor(props) {
        super(props);
        this.state = { 
            showModal: false,
         }
    }
    render() { 
        return (
            <Switch>
                {
                    compoents.map(item=> {
                        return <Route key={item.path} exact path={item.path} component={item.component}></Route>
                    })
                }
            </Switch>
        );
    }
}
 
export default ContaninerMain;