using AutoMapper;
using Microsoft.AspNetCore.Mvc;
using PokemonReviewApp.Dto;
using PokemonReviewApp.Interfaces;

namespace PokemonReviewApp.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ReviewController : Controller
    {
        private readonly IReviewRepository _ReviewRepository;
        private readonly IMapper _mapper;

        public ReviewController(IReviewRepository reviewRepository, IMapper mapper)
        {
            _ReviewRepository = reviewRepository;
            _mapper = mapper;
        }

        [HttpGet]
        public IActionResult GetReviews()
        {
            var reviews = _mapper.Map<List<ReviewDto>>(_ReviewRepository.GetReviews());

            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            return Ok(reviews);
        }

        [HttpGet("{id}")]
        public IActionResult GetReview(int id)
        {
            if (!_ReviewRepository.ReviewExists(id))
                return NotFound();

            var review = _mapper.Map<ReviewDto>(_ReviewRepository.GetReview(id));

            if (!ModelState.IsValid)
            {
                return BadRequest(ModelState);
            }

            return Ok(review);
        }

        [HttpGet("pokemon/{pokeId}")]
        public IActionResult GetReviewForAPokemon(int pokeId)
        {
            var reviews = _mapper.Map<List<ReviewDto>>(_ReviewRepository.GetReviewsOfPokemon(pokeId));

            if (!ModelState.IsValid)
                return BadRequest();

            return Ok(reviews);
        }
    }
}
