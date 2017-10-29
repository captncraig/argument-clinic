import { h, render, Component } from 'preact';

export interface HelloWorldProps {
    name: string
}

export default class HelloWorld extends Component<HelloWorldProps, any> {
    render(props) {
        return <p>Hello {props.name}!</p>
    }
}

document.addEventListener('DOMContentLoaded', function(){ 
    render(<HelloWorld name="World" />, document.querySelector('.ac-comments'));
}, false);
