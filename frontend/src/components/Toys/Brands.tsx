import './brands.css'

function Brands({visable,setBrands}){
    async function handleCheckBox(e){
        const {name,checked} = e.target
        setBrands((prev) => ({
            ...prev,
            [name]:checked
        }))
    }
    return(
        <div className={`options ${visable? "": "hideOptions"}`}>
            <label>
                    Disney
                    <input type="checkbox" name="disney" value="newsletter" onChange={handleCheckBox}/>
                    </label>
                    <label>
                    Barbie
                    <input type="checkbox" name="mattel" value="newsletter" onChange={handleCheckBox}/>
                    </label>
                    <label>
                    Lego
                    <input type="checkbox" name="lego" value="newsletter" onChange={handleCheckBox}/>
                    </label>
                    <label>
                    Hasbro
                    <input type="checkbox" name="hasbro" value="newsletter" onChange={handleCheckBox}/>
                    </label>
                    <label>
                    Hotwheels
                    <input type="checkbox" name="hotwheels" value="newsletter" onChange={handleCheckBox}/>
                    </label>
                    <label>
                    Playmates Toys
                    <input type="checkbox" name="playmatesToys" value="newsletter" onChange={handleCheckBox}/>
                    </label>
        </div>
    )
}

export default Brands
