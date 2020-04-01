import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Form, Input } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { signInProcess, changePage } from "../../actions/actions";
import { SIGNUP } from "../../pages";
import "./SignIn.css";

class SignIn extends Component {
  onFinish = values => {
    this.props.signInProcess(values.username, values.password);
  };
  changeToSignUp = () => {
    this.props.changePage(SIGNUP);
  };
  render() {
    return (
      <div className="sign-in-container">
        <Form
          name="normal_login"
          className="login-form"
          initialValues={{ remember: true }}
          onFinish={this.onFinish}
          style={{ maxWidth: 400 }}
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: "Please input your Username!" }]}
          >
            <Input
              prefix={<UserOutlined className="site-form-item-icon" />}
              placeholder="Username"
            />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: "Please input your Password!" }]}
          >
            <Input
              prefix={<LockOutlined className="site-form-item-icon" />}
              type="password"
              placeholder="Password"
            />
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              className="login-form-button"
            >
              Log in
            </Button>
            <p>
              Or <a onClick={this.changeToSignUp}>register now!</a>
            </p>
          </Form.Item>
        </Form>
      </div>
    );
  }
}
export default connect(null, { signInProcess, changePage })(SignIn);
