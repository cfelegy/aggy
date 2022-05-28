import * as React from 'react';
import { Link } from 'react-router-dom';
import { Meta, fetchAllMetas } from './api/Api';
import MetaItem from './components/MetaItem';

export default function App(): JSX.Element {
    const [metas, setMetas] = React.useState<Meta[]>([]);
    React.useEffect(() => {
        fetchAllMetas().then(res => {
            console.log(res);
            setMetas(res);
        });
    }, [])

    return (
        <main className='container'>
            <h1 className='text-center'>
                Application Mount
            </h1>
            <p>
                Here is a link to a content item: 
                <Link to="content/0">Content!</Link>
            </p>
            {metas.map((meta) => 
                <MetaItem meta={meta} />
            )}
        </main>
    );
}