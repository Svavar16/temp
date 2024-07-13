using Microsoft.EntityFrameworkCore;
using PokemonReviewApp.Models;
using System.Diagnostics.CodeAnalysis;

namespace PokemonReviewApp.Data
{
    public class DataContext : DbContext
    {
        public DataContext(DbContextOptions<DataContext> options) : base(options) { }

        public DbSet<Category> Categories { get; set; }
        public DbSet<Country> Countries { get; set; }
        public DbSet<Owner> Owners { get; set; }
        public DbSet<Pokemon> Pokemon { get; set; }
        public DbSet<PokemonOwner> PokemonsOwners { get; set; }
        public DbSet<PokemonCategory> PokemonsCategories { get; set; }
        public DbSet<Review> Reviews { get; set; }
        public DbSet<Reviewer> Reviewers { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            //base.OnModelCreating(modelBuilder);

            modelBuilder.Entity<PokemonCategory>().HasKey(pc => new { pc.PokemonId, pc.CategoryId });
            modelBuilder.Entity<PokemonCategory>().HasOne(p => p.Pokemon).WithMany(pc => pc.PokemonCategories).HasForeignKey(c => c.PokemonId);
            modelBuilder.Entity<PokemonCategory>().HasOne(p => p.Category).WithMany(pc => pc.PokemonCategories).HasForeignKey(c => c.CategoryId);

            modelBuilder.Entity<PokemonOwner>().HasKey(po => new { po.PokemonId, po.OwnerId });
            modelBuilder.Entity<PokemonOwner>().HasOne(p => p.Pokemon).WithMany(pc => pc.PokemonOwners).HasForeignKey(c => c.PokemonId);
            modelBuilder.Entity<PokemonOwner>().HasOne(p => p.Owner).WithMany(pc => pc.PokemonOwner).HasForeignKey(c => c.OwnerId);
        }
    }
}
