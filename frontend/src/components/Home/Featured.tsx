import { Link, useNavigate } from 'react-router-dom'
import './Featured.css'
import disney from '../../../images/Brand_-_Disney.webp'
import hasbro from '../../../images/hasbro.png'
import playmates from '../../../images/playmates.jpg'
import lego from '../../../images/Brand_-_Lego.webp'
import barbie from '../../../images/Brand_-_Barbie.webp'
import hotwheels from '../../../images/Brand_-_Hot_Wheels.webp'
import { filterToys } from '../../store/toys'
import { useDispatch } from 'react-redux'
import { AppDispatch } from '../../store/store'
import { useModal } from '../../Context/Modal'

function Featured(){
    const dispatch = useDispatch<AppDispatch>()
    const nav = useNavigate()
    const {brands,setObjBrands} = useModal()

    async function  handleDisney() {
        await setObjBrands({
            Disney: true,
            Hasbro: false,
            PlaymatesToys: false,
            LEGO: false,
            Mattel: false,
            Hotwheels: false
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }

    async function handleBarbie() {
        await setObjBrands({
            Disney: false,
            Hasbro: false,
            PlaymatesToys: false,
            LEGO: false,
            Mattel: true,
            Hotwheels: false
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }

    async function handlePlayMates(){
        await setObjBrands({
            Disney: false,
            Hasbro: false,
            PlaymatesToys: true,
            LEGO: false,
            Mattel: false,
            Hotwheels: false
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }
    async function handleHasbro(){
        await setObjBrands({
            Disney: false,
            Hasbro: true,
            PlaymatesToys: false,
            LEGO: false,
            Mattel: false,
            Hotwheels: false
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }

    async function handleHotwheels(){
        await setObjBrands({
            Disney: false,
            Hasbro: false,
            PlaymatesToys: false,
            LEGO: false,
            Mattel: false,
            Hotwheels: true
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }
    async function handleLego(){
        await setObjBrands({
            Disney: false,
            Hasbro: false,
            PlaymatesToys: false,
            LEGO: true,
            Mattel: false,
            Hotwheels: false
        })

        let filters = {
            theme:null,
            product: null,
            minPrice: 0,
            maxPrice:Infinity,
            brands : brands
        }
        await dispatch(filterToys(filters))
        await nav('/toys')
    }




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
                <img onClick={handleDisney} className='images' src={disney} alt='disney' />
                <img onClick={handleBarbie} className='images' src={barbie} alt='barbie' />
                <img onClick={handleLego} className='images' src={lego} alt='lego' />
                <img onClick={handleHasbro} className='images' src={hasbro} alt='hasbro' />
                <img onClick={handleHotwheels} className='images' src={hotwheels} alt='hotwheels' />
                <img onClick={handlePlayMates} className='images' src={playmates} alt='playmates' />
            </div>
        </div>
        </>
    )
}


export default Featured
