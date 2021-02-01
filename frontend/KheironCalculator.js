import React, { Component } from 'react';
import './App.css';

class KheironCalculator extends Component {
    constructor(props) {
        super(props);
        this.state = {value: '', calcResult: '', mode: 'prefix', error: ''};

        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.radioValueChange = this.radioValueChange.bind(this);
    }

    handleChange(event) {
        this.setState({value: event.target.value});
    }

    handleSubmit(event) {
        event.preventDefault();
        // Simple POST request with a JSON body using fetch
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ calculation: this.state.value })
        };
        fetch('https://4ktp77xdmh.execute-api.eu-central-1.amazonaws.com/prod/api/kheiron/' + this.state.mode, requestOptions)
            .then(response => response.json())
            .then(data => this.setState({ calcResult: data.result, error: data.error}));
    }

    radioValueChange(event) {
        this.setState({mode: event.target.value})
    }

    render() {
        return (
            <div className={"center"}>
                <h1>
                    Kheiron Calculator
                </h1>
                <div>
                    As specified in the requirements: <br/>
                    <ul>
                        <li>calculations are done in the integer domain, so <b>- / 2 10 2</b> will result in <b>-2</b> <br/></li>
                        <li>infix notation requires proper spacing between all tokens to be valid</li>
                    </ul>
                </div>
                <form onSubmit={this.handleSubmit} className={"padding"}>
                    <input type="text" value={this.state.value} onChange={this.handleChange} />
                    <input type="submit" value="Calculate" />

                </form>
                <div onChange={this.radioValueChange} className={"padding"}>
                    <input type="radio" value="prefix" name="calc-type" defaultChecked={true} /> Prefix
                    <input type="radio" value="infix" name="calc-type" /> Infix
                </div>
                <div>
                    {this.state.calcResult.length > 0 && this.state.error.length === 0 &&
                        <div>
                            <b>Result:</b> {this.state.calcResult}
                        </div>
                    }
                </div>
                <div>
                    {this.state.error.length > 0 &&
                    <div className={"warning"}>
                        <b>Error: {this.state.error}</b>
                    </div>
                    }
                </div>
            </div>
        );
    }
}

export default KheironCalculator;
