import { h, render, Component } from 'preact';

export interface HelloWorldProps {
    name: string
}

export default class HelloWorld extends Component<HelloWorldProps, any> {
    render(props) {
        return <p>Hello {props.name}!</p>
    }
}

var el = document.querySelector(".ac-comments");
if (el){
    render(<HelloWorld name="World" />, el);
}