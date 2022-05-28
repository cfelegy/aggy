const ApiBase: string = "http://localhost:4000/";

const fetch = window.fetch;

export interface Meta {
    id: string,
    provider_id: string,
    title: string,
    author: string,
    timestamp: Date,
    url: string
};

export async function fetchAllMetas(): Promise<Meta[]> {
    const res = await fetch(ApiBase, {
        method: 'GET'
    });
    if (res.status != 200) {
        // todo: communicate error message
        return [];
    }
    const body: Meta[] = await res.json();
    return body;
}