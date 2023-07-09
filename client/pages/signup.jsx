import Head from 'next/head'
import Layout from'../layout/layout'
import Link from 'next/link'
import styles from '../styles/Form.module.css'
import { useState, useContext, useEffect } from 'react'
import { API_URL } from '../constants'
import { useRouter } from 'next/router'
import { AuthContext, UserInfo } from '../modules/auth_provider'


const Index = () => {
    const [username,setUsername] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')

  const router = useRouter()


  const submitHandler = async (e) => {
    e.preventDefault()

    try {
      const res = await fetch(`${API_URL}/signup`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({username, email, password }),
      
      })
      return router.push("/login")
  
    } catch (err) {
      console.log(err)
    }
  }

    return(
            <Layout>
            <Head>
                <title>Login</title>
                </Head>
                <section className='w-3/4 mx-auto flex flex-col gap-10'>
            <div className="title">
                <h1 className='text-gray-800 text-4xl font-bold py-4'>Register</h1>
                <p className='w-3/4 mx-auto text-gray-400'>Please enter your details</p>
            </div>

            {/* form */}
            <form className='flex flex-col gap-5'>
            <div className={styles.input_group}>
                    <input 
                    type="text"
                  
                    placeholder='Username'
                    className={styles.input_text}
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                    />
                </div>
                <div className={styles.input_group}>
                    <input 
                    type="email"
                  
                    placeholder='Email'
                    className={styles.input_text}
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    />
                </div>
                <div className={styles.input_group}>
                    <input 
                    type="password"
                  
                    placeholder='password'
                    className={styles.input_text}
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    />
                </div>

                {/* login buttons */}
                <div className="input-button">
                    <button type='submit' className={styles.button} onClick={submitHandler}>
                        Signup
                    </button>
                </div>
                
                
            </form>

            {/* bottom */}
            <p className='text-center text-gray-400 '>
                 have an account? <a href='/login' className='text-purple-500' >login</a>
            </p>
        </section>

        </Layout>
    )
}
export default Index