import React from "react";
import { Form, Icon, Input, Button, message } from "antd";
import { Link } from "react-router-dom";
import { API_ROOT } from "../constants";

class NormalLoginForm extends React.Component {
  handleSubmit = e => {
    // console.log(this.props.handleLogin);
    e.preventDefault();
    this.props.form.validateFields((err, values) => {
      if (!err) {
        console.log("Received values of form: ", values);
        // fire api call
        fetch(`${API_ROOT}/login`, {
          method: "POST",
          body: JSON.stringify({
            username: values.username,
            password: values.password
          })
        })
          .then(response => {
            if (response.ok) {
              return response.text();
            }
            throw new Error(response.statusText);
          })
          .then(data => {
            message.success("Login Succeed");
            // console.log(data);
            this.props.handleLogin(data);
          })
          .catch(err => {
            message.error("Login Failed");
            console.log(err);
          });
      }
    });
  };

  render() {
    const { getFieldDecorator } = this.props.form;

    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 }
      }
    };
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0
        },
        sm: {
          span: 16,
          offset: 8
        }
      }
    };

    return (
      <Form
        {...formItemLayout}
        onSubmit={this.handleSubmit}
        className="login-form"
      >
        <Form.Item label="Username">
          {getFieldDecorator("username", {
            rules: [{ required: true, message: "Please input your username!" }]
          })(
            <Input
              prefix={<Icon type="user" style={{ color: "rgba(0,0,0,.25)" }} />}
              placeholder="Username"
            />
          )}
        </Form.Item>
        <Form.Item label="Password">
          {getFieldDecorator("password", {
            rules: [{ required: true, message: "Please input your Password!" }]
          })(
            <Input
              prefix={<Icon type="lock" style={{ color: "rgba(0,0,0,.25)" }} />}
              type="password"
              placeholder="Password"
            />
          )}
        </Form.Item>
        <Form.Item {...tailFormItemLayout}>
          <Button
            type="primary"
            htmlType="submit"
            className="login-form-button"
          >
            Log in
          </Button>
          &nbsp;&nbsp; Or <Link to="/register"> &nbsp; register now!</Link>
        </Form.Item>
      </Form>
    );
  }
}

export const Login = Form.create({ name: "LoginForm" })(NormalLoginForm);
