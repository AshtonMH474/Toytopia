import './brands.css'

function Brands({visable}){
    return(
        <div className={`options ${visable? "": "hideOptions"}`}>
            <label>
                    Disney
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
                    <label>
                    Barbie
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
                    <label>
                    Lego
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
                    <label>
                    Hasbro
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
                    <label>
                    Hotwheels
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
                    <label>
                    Playmates Toys
                    <input type="checkbox" name="subscribe" value="newsletter"/>
                    </label>
        </div>
    )
}

export default Brands
