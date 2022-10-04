import {Html, Head, Main, NextScript} from 'next/document'
import Script from "next/script";

export default function Document() {
    return (
        <Html>
            <Head>
                {/* eslint-disable-next-line @next/next/no-title-in-document-head */}
                <title>Mockxlab</title>
            </Head>
            <body>
            <Main/>
            <NextScript/>
            </body>
        </Html>
    )
}