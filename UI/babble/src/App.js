import { Layout } from "antd";
import React, { Component } from "react";
import "./App.css";
import Navbar from "./components/Navbar/Navbar";
import Feed from "./components/Feed/Feed";
import SignIn from "./components/SignIn/SignIn";
import SignUp from "./components/SignUp/SignUp";
import { connect } from "react-redux";
import { FEED, LOGIN, SIGNUP, USER, SUBSCRIPTION } from "./pages";
import UserPage from "./components/UserPage/UserPage";
import Subscription from "./components/Subscription/Subscription";

const { Footer, Content } = Layout;

class App extends Component {
  renderPage = (page) => {
    switch (page) {
      case LOGIN:
        return (<SignIn />);
      case SIGNUP:
        return (<SignUp />);
      case FEED:
        return (<Feed />);
      case USER:
        return (<UserPage />);
      case SUBSCRIPTION:
        return (<Subscription />);
      default:
        return (<SignIn />);
    }
  }
  render() {
    return (
      <Layout className="App">
        <Navbar />
        <Content
          className="site-layout"
          style={{
            padding: "15px 50px",
            marginTop: 64,
            height: "100%",
            overflow: "initial"
          }}
        >
          <div
            className="site-layout-background"
            style={{
              paddingLeft: 24,
              paddingRight: 24,
              marginTop: "15px",
              height: "95%",
              paddingTop: "25px",
              paddingBottom: "15px"
            }}
          >
            {this.renderPage(this.props.page)}
          </div>
        </Content>
        <Footer
          style={{
            textAlign: "center",
            bottom: 0,
            width: "100%"
          }}
        >
          Babble ©2020
        </Footer>
      </Layout>
    );
  }
}

const mapStateToProps = state => ({
  page: state.page
});

export default connect(mapStateToProps, {})(App);
