using Microsoft.EntityFrameworkCore;
using superHeroAPI.Models;

namespace superHeroAPI.Data
{
    public class DataContext : DbContext
    {
        public DataContext(DbContextOptions<DataContext> options) : base(options) { }

        public DbSet<SuperHero> superHeroes { get; set; }
    }
}
