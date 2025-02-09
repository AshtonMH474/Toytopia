import { Link } from 'react-router-dom'
import './Featured.css'

function Featured(){
    return (
        <>
        <div>
            <div className='info'>
                <h1>Shop By Brand</h1>
                <Link to='/toys' className='link'>
                Shop All Brands
                </Link>
            </div>
        </div>
        </>
    )
}


export default Featured
