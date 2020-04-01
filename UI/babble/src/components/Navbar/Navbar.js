import { Avatar, Dropdown, Layout, Menu, Typography } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { changePage, loadSubscriptionDetails, loadUserDetails, signOutProcess } from "../../actions/actions";
import { FEED, SUBSCRIPTION } from "../../pages";
import "./Navbar.css";

const { Header } = Layout;
const { Title } = Typography;

class Navbar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      selectedKeys: ["1"]
    };
  }

  goToLogin = () => {
    this.props.signOutProcess(this.props.username);
  };
  goToFeed = () => {
    this.props.changePage(FEED);
    this.setState({ selectedKeys: ["1"] });
  };
  goToMyPage = () => {
    this.props.loadUserDetails(this.props.token, this.props.username);
    this.setState({ selectedKeys: [] });
  };
  goToSubscriptions = () => {
    this.props.changePage(SUBSCRIPTION);
    this.setState({ selectedKeys: ["2"] });
  };
  menu = (
    <Menu theme="dark">
      <Menu.Item key="1" onClick={this.goToMyPage}>
        My Page
      </Menu.Item>
      <Menu.Item key="2" onClick={this.goToLogin}>
        Sign Out
      </Menu.Item>
    </Menu>
  );
  conditionallyRenderMenu = token => {
    if (token !== "") {
      return (
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={["1"]}
          selectedKeys={this.state.selectedKeys}
          style={{ float: "right", marginRight: 24 }}
        >
          <Menu.Item key="1" onClick={this.goToFeed}>
            Feed
          </Menu.Item>
          <Menu.Item
            key="2"
            onClick={this.goToSubscriptions}
            style={{ marginRight: 25 }}
          >
            Subscriptions
          </Menu.Item>
          <Dropdown overlay={this.menu} placement="bottomCenter">
            <Avatar>{this.props.username[0].toUpperCase()}</Avatar>
          </Dropdown>
        </Menu>
      );
    }
  };
  render() {
    return (
      <Header style={{ position: "fixed", zIndex: 1, width: "100%" }}>
        <div className="logo">
          <Title style={{ color: "#fff" }}>Babble</Title>
        </div>
        {this.conditionallyRenderMenu(this.props.token)}
      </Header>
    );
  }
}

const mapStateToProps = state => ({
  username: state.me.username,
  token: state.token,
  page: state.page
});

export default connect(mapStateToProps, {
  signOutProcess,
  loadUserDetails,
  changePage,
  loadSubscriptionDetails
})(Navbar);
