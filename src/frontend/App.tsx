import * as React from 'react';
import { Link } from 'react-router-dom';

export default function App(): JSX.Element {
    return (
        <main className='container'>
                <h1 className='text-center'>
                    Application Mount
                </h1>
                <p>
                    Here is a link to a content item: 
                    <Link to="content/0">Content!</Link>
                </p>
            </main>
    );
}