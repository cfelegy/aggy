import * as React from 'react';

class App extends React.Component<IAppProps, IAppState> {
    render() {
        return (
            <main className='container'>
                <h1 className='text-center'>
                    Application Mount
                </h1>
            </main>
        )
    }
}

export interface IAppProps {}
export interface IAppState {}

export default App;