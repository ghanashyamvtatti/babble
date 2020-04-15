import { Button, Form, Input } from "antd";
import React, { Component } from "react";
import { connect } from "react-redux";
import { signUpProcess } from "../../actions/actions";
import "./SignUp.css";

class SignUp extends Component {
  onFinish = values => {
    this.props.signUpProcess(values.name, values.username, values.password);
  };
  render() {
    return (
      <div className="sign-up-container">
        <Form
          name="normal_signup"
          className="signup-form"
          initialValues={{ remember: true }}
          onFinish={this.onFinish}
          style={{ maxWidth: 400 }}
        >
          <Form.Item
            name="name"
            rules={[
              { required: true, message: "Please input your Full Name!" }
            ]}
          >
            <Input placeholder="Full Name" />
          </Form.Item>
          <Form.Item
            name="username"
            rules={[{ required: true, message: "Please input your Username!" }]}
          >
            <Input placeholder="Username" />
          </Form.Item>
          <Form.Item
            name="password"
            rules={[{ required: true, message: "Please input your Password!" }]}
          >
            <Input type="password" placeholder="Password" />
          </Form.Item>

          <Form.Item>
            <Button
              type="primary"
              htmlType="submit"
              className="login-form-button"
            >
              Sign Up
            </Button>
          </Form.Item>
        </Form>
      </div>
    );
  }
}

export default connect(null, { signUpProcess })(SignUp);
