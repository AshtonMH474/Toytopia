import { Link } from 'react-router-dom'
import './Featured.css'
import disney from '../../../images/Brand_-_Disney.webp'
import hasbro from '../../../images/hasbro.png'
import playmates from '../../../images/playmates.jpg'
import lego from '../../../images/Brand_-_Lego.webp'
import barbie from '../../../images/Brand_-_Barbie.webp'
import hotwheels from '../../../images/Brand_-_Hot_Wheels.webp'

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
            <div className='brands'>
                <img className='images' src={disney} alt='disney' />
                <img className='images' src={barbie} alt='barbie' />
                <img className='images' src={lego} alt='lego' />
                <img className='images' src={hasbro} alt='hasbro' />
                <img className='images' src={hotwheels} alt='hotwheels' />
                <img className='images' src={playmates} alt='playmates' />
            </div>
        </div>
        </>
    )
}


export default Featured
