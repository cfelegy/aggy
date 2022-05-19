import * as React from 'react';
import { Link, useParams } from 'react-router-dom';

export default function Content(): JSX.Element {
    let params = useParams();
    
    return (
        <main className='container'>
            <h1>Content</h1>
            <h2>ID: {params.id}</h2>
            <Link to="/">Go back home</Link>
        </main>
    );
}