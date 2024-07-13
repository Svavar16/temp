using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using PokemonReviewApp.Dto;
using PokemonReviewApp.Interfaces;
using PokemonReviewApp.Models;

namespace PokemonReviewApp.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class OwnerController : ControllerBase
    {
        private readonly IOwnerRepository _OwnerRepositry;
        private readonly ICountryRepository _CountryRepositry;
        private readonly IMapper _Mapper;
        public OwnerController(IOwnerRepository ownerRepositry, IMapper mapper, ICountryRepository countryRepository)
        {
            _OwnerRepositry = ownerRepositry;
            _Mapper = mapper;
            _CountryRepositry = countryRepository;
        }

        [HttpGet]
        public IActionResult GetOwners()
        {
            var owners = _Mapper.Map<List<OwnerDto>>(_OwnerRepositry.GetOwners()).ToList();

            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            return Ok(owners);
        }

        [HttpGet("{ownerId}")]
        public IActionResult GetOwner(int ownerId)
        {
            if (!_OwnerRepositry.OwnerExists(ownerId))
            {
                return NotFound();
            }

            var owner = _Mapper.Map<OwnerDto>(_OwnerRepositry.GetOwner(ownerId));

            if (!ModelState.IsValid)
                return BadRequest();

            return Ok(owner);
        }

        [HttpGet("{ownerId}/pokemon")]
        public IActionResult GetPokemonByOwner(int ownerId)
        {
            if (!_OwnerRepositry.OwnerExists(ownerId))
                return NotFound();

            var owner = _Mapper.Map<List<PokemonDto>>(_OwnerRepositry.GetPokemonByOwner(ownerId)).ToList();

            if (!ModelState.IsValid)
                return BadRequest();

            return Ok(owner);
        }



        [HttpPost]
        public IActionResult PostOwner([FromQuery] int countryId, OwnerDto ownerCreate)
        {
            if (ownerCreate == null) return BadRequest(ModelState);

            var owners = _OwnerRepositry.GetOwners().Where(c => c.LastName.Trim().ToUpper() == ownerCreate.LastName.Trim().ToUpper()).FirstOrDefault();

            if (owners == null)
            {
                return BadRequest("Owner already exists");
            }

            if (!ModelState.IsValid) return BadRequest(ModelState);

            var ownerMap = _Mapper.Map<Owner>(ownerCreate);

            ownerMap.Country = _CountryRepositry.GetCountry(countryId);

            if(!_OwnerRepositry.create)

            return Ok();
        }
    }
}
