"use client"
import styles from './page.module.css'
import {useState} from "react";

export default function Home() {
    const [foo, setFoo] = useState('')
    async function  getFoo(){
        const r = await fetch('http://localhost:3001/')
        const json = await r.json()
        setFoo(json['foo'])
    }
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <p>
          foo: {foo}
        </p>
          <button onClick={getFoo}>get foo</button>
      </div>

    </main>
  )
}
