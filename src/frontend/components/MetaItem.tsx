import * as React from 'react';
import { Meta } from '../api/Api';

type MetaItemProps = {
    meta: Meta
};

export default function MetaItem(props: MetaItemProps): JSX.Element {
    const meta = props.meta;

    return (<>
        <p>{meta.title}</p>
        <p>{meta.author}</p>
        <p>{meta.url}</p>
    </>);
}