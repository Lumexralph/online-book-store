import React from 'react';
import { Form, Input, Button, Checkbox } from 'antd';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

import { ProductServiceClient } from "./services/StoreServiceClientPb";
import { Timestamp, Product, AddProductRequest } from './services/store_pb';

const productService = new ProductServiceClient("http://localhost:8080");

let today = new Date()
const currentTimestamp = new google_protobuf_timestamp_pb.Timestamp();
currentTimestamp.fromDate(today);
const time = new Timestamp();
time.setTimestamp(currentTimestamp);

const newProduct = new Product();
newProduct.setId(2);
newProduct.setUuid("12ht6-78");
newProduct.setName("Sherlock Holmes Jnr");
newProduct.setDescription("It is an investigative book: sequel");
newProduct.setPrice(153.56);
newProduct.setSlug("sherlock-holmes-jnr");
newProduct.setQuantity(50);
newProduct.setCreatedAt(time);

const request = new AddProductRequest();
request.setProduct(newProduct);

// create the product
productService.addProduct(request, {})
.then(data => console.log("product created: ", data))
.catch(err => {
  console.log("could not create Product");
  console.log(err);
});

const layout = {
  labelCol: { span: 8 },
  wrapperCol: { span: 16 },
};
const tailLayout = {
  wrapperCol: { offset: 8, span: 16 },
};

const ProductClient = () => {
  const onFinish = (values: any) => {
    console.log('Success:', values);
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log('Failed:', errorInfo);
  };

  return (
    <Form
      {...layout}
      name="basic"
      initialValues={{ remember: true }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
    >
      <Form.Item
        label="Username"
        name="username"
        rules={[{ required: true, message: 'Please input your username!' }]}
      >
        <Input />
      </Form.Item>

      <Form.Item
        label="Password"
        name="password"
        rules={[{ required: true, message: 'Please input your password!' }]}
      >
        <Input.Password />
      </Form.Item>

      <Form.Item {...tailLayout} name="remember" valuePropName="checked">
        <Checkbox>Remember me</Checkbox>
      </Form.Item>

      <Form.Item {...tailLayout}>
        <Button type="primary" htmlType="submit">
          Submit
        </Button>
      </Form.Item>
    </Form>
  );
};

export default ProductClient;
