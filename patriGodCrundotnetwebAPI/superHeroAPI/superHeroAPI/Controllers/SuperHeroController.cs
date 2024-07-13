using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using superHeroAPI.Data;
using superHeroAPI.Models;

namespace superHeroAPI.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class SuperHeroController : ControllerBase
    {
        private readonly DataContext _dataContext;

        public SuperHeroController(DataContext dataContext)
        {
            _dataContext = dataContext;
        }

        private static List<SuperHero> heroes = new List<SuperHero>
        {
            new SuperHero { Id = 1, Name = "Batman", FirstName = "Bruce", LastName = "Wayne", Place = "Gotham" }
        };
        [HttpGet]
        public async Task<ActionResult<List<SuperHero>>> Get()
        {
            return Ok(await _dataContext.superHeroes.ToListAsync<SuperHero>());
        }

        [HttpGet("{id}")]
        public async Task<ActionResult<SuperHero>> GetSingle(int id)
        {
            
            var hero = await _dataContext.superHeroes.FindAsync(id);
            if (hero == null)
            {
                return BadRequest("Hero not found.");
            }
            return Ok(hero);
        }

        [HttpPost]
        public async Task<ActionResult<List<SuperHero>>> AddHero(SuperHero hero)
        {
            _dataContext.superHeroes.Add(hero);
            await _dataContext.SaveChangesAsync();

            return Ok(await _dataContext.superHeroes.ToListAsync<SuperHero>());
        }

        [HttpPut]
        public async Task<ActionResult<List<SuperHero>>> UpdateHero(SuperHero request)
        {
            var hero = await _dataContext.superHeroes.FindAsync(request.Id);
            if (hero == null)
            {
                return BadRequest("Hero not found.");
            }
            hero.Name = request.Name;
            hero.FirstName = request.FirstName;
            hero.LastName = request.LastName;
            hero.Place = request.Place;

            return Ok(await _dataContext.superHeroes.ToListAsync<SuperHero>());
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult<List<SuperHero>>> DeleteHero(int id)
        {
            var hero = await _dataContext.superHeroes.FindAsync(id);
            if (hero == null)
            {
                return BadRequest("Hero not found");
            }
            _dataContext.superHeroes.Remove(hero);
            await _dataContext.SaveChangesAsync();

            return Ok(await _dataContext.superHeroes.ToListAsync<SuperHero>());
        }
    }
}
