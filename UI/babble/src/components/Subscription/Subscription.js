import { DeleteTwoTone, PlusCircleOutlined } from "@ant-design/icons";
import { Avatar, Button, List } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { getAllUsers, loadSubscriptionDetails, subscribe, unsubscribe } from "../../actions/actions";

class Subscription extends Component {
  componentDidMount() {
    this.props.loadSubscriptionDetails(this.props.token, this.props.username);
    this.props.getAllUsers(this.props.token);
  };
  subscribeUser = (subscriber, publisher) => {
    this.props.subscribe(this.props.token, subscriber, publisher);
  }
  unsubscribeUser = (subscriber, publisher) => {
    this.props.unsubscribe(this.props.token, subscriber, publisher);
  };
  goToProfile = username => {
    this.props.loadUserDetails(this.props.token, username);
  };
  render() {
    return (
      <div>
        <List
          bordered
          dataSource={this.props.users}
          renderItem={(user, index) => (
            <List.Item key={index}>
              <List.Item.Meta
                avatar={<Avatar>{user.username[0].toUpperCase()}</Avatar>}
                title={user.name}
              />
              {this.props.subscriptions.includes(user.username) ? (
                <Button
                  icon={<DeleteTwoTone />}
                  onClick={() =>
                    this.unsubscribeUser(this.props.username, user.username)
                  }
                />
              ) : (
                <Button
                  icon={<PlusCircleOutlined />}
                  onClick={() =>
                    this.subscribeUser(this.props.username, user.username)
                  }
                />
              )}
            </List.Item>
          )}
        />
      </div>
    );
  }
}

const mapStateToProps = state => ({
  username: state.me.username,
  token: state.token,
  subscriptions: state.subscriptions,
  users: state.data
});

export default connect(mapStateToProps, {
  getAllUsers,
  loadSubscriptionDetails,
  unsubscribe,
  subscribe
})(Subscription);
